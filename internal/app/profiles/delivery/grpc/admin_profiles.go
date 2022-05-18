package grpc

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/profiles"
	proto "2021_1_Noskool_team/internal/app/profiles/delivery/grpc/protobuff"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type adminProfiles struct {
	proto.UnimplementedAdminProfilesServer
	adminProfiles  profiles.Usecase
	sessionsClient client.AuthCheckerClient
}

func newAdminProfilesSrv(config *configs.Config, grpcServer *grpc.Server, profiles profiles.Usecase) {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	adminProfiles := &adminProfiles{
		adminProfiles:  profiles,
		sessionsClient: client.NewSessionsClient(grpcCon),
	}

	proto.RegisterAdminProfilesServer(grpcServer, adminProfiles)
	reflection.Register(grpcServer)
}

func StartSessionsGRPCServer(url string, config *configs.Config, profiles profiles.Usecase) {
	logrus.Info(url)
	list, err := net.Listen("tcp", url)
	if err != nil {
		logrus.Error(err)
	}

	grpcSrv := grpc.NewServer()

	newAdminProfilesSrv(config, grpcSrv, profiles)

	err = grpcSrv.Serve(list)
	if err != nil {
		logrus.Error(err)
	}
}

func (a adminProfiles) CreateAdminProfile(ctx context.Context, in *proto.CreateAdminProfileIn) (*proto.CreateAdminProfileOut, error) {
	adminProfile := &models.AdminProfile{
		Login:             in.Profile.GetLogin(),
		EncryptedPassword: in.Profile.GetPassword(),
	}

	err := a.adminProfiles.CreateAdminProfile(adminProfile)
	if err != nil {
		out := &proto.CreateAdminProfileOut{
			Error: proto.CreateAdminProfileOut_ERROR_CREATING_PROFILE,
		}
		return out, err
	}

	result, err := a.sessionsClient.CreateAdminSession(ctx, in.GetProfile().GetLogin())
	if err != nil {
		out := &proto.CreateAdminProfileOut{
			Error: proto.CreateAdminProfileOut_ERROR_CREATING_PROFILE,
		}
		return out, err
	}

	fmt.Println(result.Hash)
	out := &proto.CreateAdminProfileOut{
		Error:   proto.CreateAdminProfileOut_NO_ERROR,
		Session: result.Hash,
	}

	return out, nil
}

func (a adminProfiles) LoginAdminProfile(ctx context.Context, in *proto.LoginAdminProfileIn) (*proto.LoginAdminProfileInOut, error) {
	adminProfile := &models.AdminProfile{
		Login:             in.Profile.GetLogin(),
		EncryptedPassword: in.Profile.GetPassword(),
	}

	err := a.adminProfiles.LoginAdminProfile(adminProfile)
	if err != nil {
		out := &proto.LoginAdminProfileInOut{
			Error: proto.LoginAdminProfileInOut_ERROR_LOGIN_PROFILE,
		}
		return out, err
	}

	result, err := a.sessionsClient.CreateAdminSession(ctx, in.GetProfile().GetLogin())
	if err != nil {
		out := &proto.LoginAdminProfileInOut{
			Error: proto.LoginAdminProfileInOut_ERROR_LOGIN_PROFILE,
		}
		return out, err
	}

	fmt.Println(result.Hash)
	out := &proto.LoginAdminProfileInOut{
		Error:   proto.LoginAdminProfileInOut_ERROR_LOGIN_PROFILE,
		Session: result.Hash,
	}

	return out, nil

}

func (a adminProfiles) CheckIsLogged(ctx context.Context, in *proto.CheckIsLoggedIn) (*proto.CheckIsLoggedOut, error) {
	_, err := a.sessionsClient.Check(ctx, in.GetSession())
	if err != nil {
		return nil, err
	}

	return &proto.CheckIsLoggedOut{IsLogged: true}, nil
}

func (a adminProfiles) LogoutAdminProfile(ctx context.Context, in *proto.LoginAdminProfileIn) (*proto.LoginAdminProfileInOut, error) {
	panic("implement me")
}
