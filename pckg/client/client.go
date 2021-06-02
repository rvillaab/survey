package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	pb "survey/pckg/question"
	"time"

	"google.golang.org/grpc"
)

var (
	serverAddr         = flag.String("server_addr", ":10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {

	var optionReceived = ""
	var parameter = ""

	if len(os.Args) > 1 {
		optionReceived = os.Args[1]
	}

	if len(os.Args) > 2 {
		parameter = os.Args[2]
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTimeout(10*time.Second))

	conn, err := grpc.Dial(*serverAddr, opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	processRequest(optionReceived, parameter, conn, ctx)

	defer cancel()

}

func processRequest(option string, parameter string, conn *grpc.ClientConn, ctx context.Context) {

	client := pb.NewQuestionServiceClient(conn)

	switch option {
	case "1": //Create question
		var newQuestion = &pb.Question{}
		jsonErr := json.Unmarshal([]byte(parameter), &newQuestion)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		fmt.Print(fmt.Sprintf("content: %v, Desc:%v, anser:%v \n", newQuestion.Content, newQuestion.Description, newQuestion.Answer))

		response, err := client.CreateQuestion(ctx, newQuestion)

		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
		break
	case "2": //Get All question
		response, err := client.GetAllQuestions(ctx, &pb.EmptyRequest{})

		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
		break
	case "3": //Update question
		var questionUpdate = &pb.Question{}
		jsonErr := json.Unmarshal([]byte(parameter), &questionUpdate)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		response, err := client.UpdateQuestion(ctx, questionUpdate)

		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
		break
	case "4": //Delete question
		response, err := client.DeleteQuestion(ctx, &pb.RequestWithId{Id: parameter})

		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
		break
	case "5": //Get question by Id
		response, err := client.GetQuestionById(ctx, &pb.RequestWithId{Id: parameter})

		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response)
		break
	default:
		fmt.Println("Not a valid option")
	}

}
