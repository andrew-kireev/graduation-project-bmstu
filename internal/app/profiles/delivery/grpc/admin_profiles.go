package grpc

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/profiles"
	proto "2021_1_Noskool_team/internal/app/profiles/delivery/grpc/protobuff"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
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

	out := &proto.CreateAdminProfileOut{
		Error: proto.CreateAdminProfileOut_NO_ERROR,
	}

	return out, nil
}

func (a adminProfiles) LoginAdminProfile(ctx context.Context, in *proto.LoginAdminProfileIn) (*proto.LoginAdminProfileInOut, error) {
	panic("implement me")
}

func (a adminProfiles) LogoutAdminProfile(ctx context.Context, in *proto.LoginAdminProfileIn) (*proto.LoginAdminProfileInOut, error) {
	panic("implement me")
}
