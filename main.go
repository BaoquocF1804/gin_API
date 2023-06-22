package main

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/gapi"
	"github.com/techschool/simplebank/pb"
	"github.com/techschool/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

var conn *sql.DB

const serverAddress = "0.0.0.0:8080"

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	cfg := mysql.Config{
		User:                 ("root"),
		Passwd:               ("secret"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "simplebank",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	// Get a database handle.
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	//server := api.NewSever(store)
	//err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store *db.Store) {
	server, err := gapi.NewSever(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}
	log.Printf("start gRPC %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC: ", err)
	}
}

func runGatewayServer(config util.Config, store *db.Store) {
	server, err := gapi.NewSever(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler sever", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}
	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP gateway server: ", err)
	}
}
