package mock

import (
	"github.com/golang/mock/gomock"
	"user-service/internal/user"
)

func MockRepositoryIsEmailExistsTrue(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsEmailExists(MockUser.Email).Return(true, nil).Times(1)
}

func MockRepositoryIsEmailExistsFalse(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsEmailExists(MockUser.Email).Return(false, nil).Times(1)
}

func MockRepositoryIsNicknameExistsTrue(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsNicknameExists(MockUser.Nickname).Return(true, nil).Times(1)
}

func MockRepositoryIsNicknameExistsFalse(mockRepository MockIRepository) {
	mockRepository.EXPECT().IsNicknameExists(MockUser.Nickname).Return(false, nil).Times(1)
}

func MockRepositoryUpsert(mockRepository MockIRepository) {
	mockRepository.EXPECT().Upsert(gomock.Any()).Return(nil).Times(1)
}

func MockRepositoryGetUsers(mockRepository MockIRepository) {
	mockRepository.EXPECT().GetUsers().Return(MockUsers, nil).Times(1)
}

func MockRepositoryGetUser(mockRepository MockIRepository) {
	mockRepository.EXPECT().GetUserByID(MockUser.ID).Return(MockUser, nil).Times(1)
}

func MockRepositoryGetUserNotFound(mockRepository MockIRepository) {
	mockRepository.EXPECT().GetUserByID(MockUser.ID).Return(new(user.User), nil).Times(1)
}

func MockRepositoryDeleteUser(mockRepository MockIRepository) {
	mockRepository.EXPECT().DeleteUserByID(MockUser.ID).Return(nil).Times(1)
}

func MockRepositoryDeleteUserNotFound(mockRepository MockIRepository) {
	mockRepository.EXPECT().DeleteUserByID(MockUser.ID).Return(nil).Times(1)
}
