package mocks

import "restAPI/pkg/models"

var pic = "VGhpcyBpcyBhIHRlc3QgaW1hZ2U"

var User = []models.User{
	{Username: "test", Email: "test@gmail.com", Uid: "x5ATVabOCvYhlKDhKFw71tn2zIU2", Picture: &pic, Following: []string{"Joe"}, Friends: []string{"Bob"}},
	{Username: "test2", Email: "test2@gmail.com", Uid: "z5ATVabOCvYhlKDhKFw71tn2zIU2", Picture: nil, Following: []string{}, Friends: []string{}},
}
