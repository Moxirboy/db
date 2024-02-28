package handler
import (pb "student/proto/exam"
	"student/internal/models"
)


func FromProtoToResponse(
	res []models.QuestionsAll,
)  *pb.TestResponse {
	questions:=[]*pb.Test{}
	for _,question:=range res{
		questions=append(questions,FromProtoResponseToTest(question))
	}
	return &pb.TestResponse{
        Test:questions,
    }
}



func FromProtoResponseToTest(
	res models.QuestionsAll,
) *pb.Test {
	choices:=[]*pb.Choices{}
	for _,choice:=range res.Choices{
		choices=append(choices,FromChoiceProtoToDto(choice))
	}
	return &pb.Test{
		Id:                res.Id,
		Question:          res.QuestionText,
		Choices:           choices,
}}

func FromChoiceProtoToDto(
	res models.Choices,
)   *pb.Choices{
	return &pb.Choices{
		Id:  res.Id,
		Choices: res.ChoiceText,
		IsCorrect: res.Is_correct,
}}
