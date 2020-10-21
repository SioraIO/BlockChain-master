type node struct {
}

func main() {

	node := node{}
	node.Start()

	select {}
}

func (n *node) Start() {
	go n.Listen()
	go n.Dial()
}

func (n *node) Listen() {
	listen, err := net.Listen("tcp", ":8222")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pEcho := peer.PeerEcho{}

	grpcServer := grpc.NewServer()

	peer.RegisterPeerEchoServer(grpcServer, &pEcho)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to star %v", err)
	}
}

func (n *node) Dial() {
	conn, err := grpc.Dial(":8222", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := peer.NewPeerEchoClient(conn)

	res, err := client.Request(context.Background(), &peer.PeerRequest{Id: "007", Msg: "hello"})
	if err != nil {
		log.Fatalf("failed to call broadcast: %v", err)
	}

	log.Printf("%s", res.Msg)
}
