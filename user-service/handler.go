package main

import (
	"log"
	"errors"
	"context"
	pb "shippy/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	repo Repository
	tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPasswd)
	if err:=h.repo.Create(req); err != nil {
		return nil
	}
	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	u, err := h.repo.Get(req.Id);
	if err != nil {
		return err
	}
	resp.User = u
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	plainPasswd := req.Password
	u, err := h.repo.GetByEmailAndPassword(req)
	if err != nil {
		log.Println("[x]user-service[handler] Error get user; Errorinfo: ", err)
		return err
	}
	log.Println("[!]user-service[handler] get user : ", u)
	log.Println("[!]user-service[handler] req user : ", req)
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPasswd)); err != nil {
		log.Println("[x]user-service[handler] hashed password comparision error; Errorinfo: ", err)
		return err
	}
	t, err := h.tokenService.Encode(u)
	if err != nil {
		return err
	}
	resp.Token = t
	log.Println("[!]user-service[handler] Token= ", t)
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return nil
	}
	if claims.User.Id == "" {
		return errors.New("invalid user")
	}
	resp.Valid = true
	return nil
}