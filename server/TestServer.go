package server

import (
	"context"
	"fmt"
	"io"
	"log"

	"platzi.com/go/grpc/models"
	"platzi.com/go/grpc/repository"
	"platzi.com/go/grpc/studentpb"
	"platzi.com/go/grpc/testpb"
)

type TestServer struct {
	repo repository.Repository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}

	err := s.repo.SetTest(ctx, test)

	if err != nil {
		return nil, err
	}

	return &testpb.SetTestResponse{
		Id: test.Id,
	}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		question := &models.Question{
			Id:       msg.GetId(),
			TestId:   msg.GetTestId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
		}

		err = s.repo.SetQuestion(context.Background(), question)

		if err != nil {
			fmt.Println(err)
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		enrollment := &models.Enrollment{
			TestId:    msg.GetTestId(),
			StudentId: msg.GetStudentId(),
		}

		err = s.repo.SetEnrollment(context.Background(), enrollment)

		if err != nil {
			fmt.Println(err)
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) GetStudentsPerTest(req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := s.repo.GetStudentsPerTest(context.Background(), req.GetTestId())

	if err != nil {
		return err
	}

	for key := 0; key < len(students); key++ {
		repoStudent := students[key]
		student := &studentpb.Student{
			Id:   repoStudent.Id,
			Name: repoStudent.Name,
			Age:  repoStudent.Age,
		}

		err := stream.Send(student)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		questions, err := s.repo.GetQuestionPerTest(context.Background(), msg.GetTestId())

		if err != nil {
			return err
		}

		var currentQuestion = &models.Question{}
		i := 0
		lenQuestions := len(questions)
		lenQuestions32 := int32(lenQuestions)

		for {
			if i < lenQuestions {
				currentQuestion = questions[i]
				questionToSend := &testpb.QuestionPerTest{
					Id:       currentQuestion.Id,
					Question: currentQuestion.Question,
					Ok:       false,
					Current:  int32(i + 1),
					Total:    lenQuestions32,
				}

				err := stream.Send(questionToSend)

				if err != nil {
					return err
				}

				i++

				answer, err := stream.Recv()

				if err == io.EOF {
					return nil
				}

				if err != nil {
					return err
				}

				log.Println("Answer: ", answer.GetAnswer())
				currentQuestion.Answer = answer.GetAnswer()

				err = s.repo.SetAnswer(context.Background(), currentQuestion)

				if err != nil {
					fmt.Println(err)
					return err
				}
			} else {
				questionToSend := &testpb.QuestionPerTest{
					Id:       "",
					Question: "",
					Ok:       true,
					Current:  int32(0),
					Total:    int32(0),
				}

				err := stream.Send(questionToSend)

				if err != nil {
					return err
				}

				break
			}
		}
	}
}
