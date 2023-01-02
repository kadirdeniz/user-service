package mock

import (
	"github.com/golang/mock/gomock"
	"user-service/internal/user"
)

func MockRepositoryIsEmailExistsTrue(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsEmailExists(user.MockUser.Email).Return(true, nil).Times(1)
}

func MockRepositoryIsEmailExistsFalse(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsEmailExists(user.MockUser.Email).Return(false, nil).Times(1)
}

func MockRepositoryIsNicknameExistsTrue(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsNicknameExists(user.MockUser.Nickname).Return(true, nil).Times(1)
}

func MockRepositoryIsNicknameExistsFalse(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsNicknameExists(user.MockUser.Nickname).Return(false, nil).Times(1)
}

func MockRepositoryUpsert(mockRepository MockIRepository) {
	mockRepository.EXPECT().Upsert(gomock.Any()).Return(nil).Times(1)
}

func MockRepositoryGetUsers(mockRepository MockIRepository) {
	mockRepository.EXPECT().GetUsers().Return(user.MockUsers, nil).Times(1)
}

func MockRepositoryGetUser(mockRepository MockIRepository) {
	mockRepository.EXPECT().GetUserByID(user.MockUser.ID).Return(user.MockUser, nil).Times(1)
}

func MockRepositoryGetUserNotFound(mockRepository MockIRepository) {
	mockRepository.EXPECT().GetUserByID(user.MockUser.ID).Return(new(user.User), nil).Times(1)
}

func MockRepositoryDeleteUser(mockRepository MockIRepository) {
	mockRepository.EXPECT().DeleteUserByID(user.MockUser.ID).Return(nil).Times(1)
}

func MockRepositoryDeleteUserNotFound(mockRepository MockIRepository) {
	mockRepository.EXPECT().DeleteUserByID(user.MockUser.ID).Return(nil).Times(1)
}
