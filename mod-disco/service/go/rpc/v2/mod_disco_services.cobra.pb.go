// Code generated by protoc-gen-cobra. DO NOT EDIT.

package v2

import (
	client "github.com/getcouragenow/protoc-gen-cobra/client"
	flag "github.com/getcouragenow/protoc-gen-cobra/flag"
	iocodec "github.com/getcouragenow/protoc-gen-cobra/iocodec"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func SurveyServiceClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("SurveyService"),
		Short: "SurveyService service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_SurveyServiceNewSurveyProjectCommand(cfg),
		_SurveyServiceGetSurveyProjectCommand(cfg),
		_SurveyServiceListSurveyProjectCommand(cfg),
		_SurveyServiceUpdateSurveyProjectCommand(cfg),
		_SurveyServiceDeleteSurveyProjectCommand(cfg),
		_SurveyServiceNewSurveyUserCommand(cfg),
		_SurveyServiceGetSurveyUserCommand(cfg),
		_SurveyServiceListSurveyUserCommand(cfg),
		_SurveyServiceUpdateSurveyUserCommand(cfg),
		_SurveyServiceDeleteSurveyUserCommand(cfg),
		_SurveyServiceNewDiscoProjectCommand(cfg),
		_SurveyServiceGetDiscoProjectCommand(cfg),
		_SurveyServiceListDiscoProjectCommand(cfg),
		_SurveyServiceUpdateDiscoProjectCommand(cfg),
		_SurveyServiceDeleteDiscoProjectCommand(cfg),
	)
	return cmd
}

func _SurveyServiceNewSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &NewSurveyProjectRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("NewSurveyProject"),
		Short: "NewSurveyProject RPC client",
		Long:  "Projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "NewSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &NewSurveyProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SysAccountProjectRefId, cfg.FlagNamer("SysAccountProjectRefId"), "", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.SupportRoleTypes, cfg.FlagNamer("SupportRoleTypes"), "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.UserNeedTypes, cfg.FlagNamer("UserNeedTypes"), "")

	return cmd
}

func _SurveyServiceGetSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetSurveyProject"),
		Short: "GetSurveyProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GetSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceListSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &ListRequest{
		IdRequest: &IdRequest{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ListSurveyProject"),
		Short: "ListSurveyProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "ListSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &ListRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyUserId, cfg.FlagNamer("IdRequest SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountProjectId, cfg.FlagNamer("IdRequest SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyProjectId, cfg.FlagNamer("IdRequest SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountAccountId, cfg.FlagNamer("IdRequest SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.DiscoProjectId, cfg.FlagNamer("IdRequest DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountOrgId, cfg.FlagNamer("IdRequest SysAccountOrgId"), "", "")
	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")

	return cmd
}

func _SurveyServiceUpdateSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &UpdateSurveyProjectRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("UpdateSurveyProject"),
		Short: "UpdateSurveyProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "UpdateSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &UpdateSurveyProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UpdateSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.SupportRoleTypes, cfg.FlagNamer("SupportRoleTypes"), "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.UserNeedTypes, cfg.FlagNamer("UserNeedTypes"), "")

	return cmd
}

func _SurveyServiceDeleteSurveyProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("DeleteSurveyProject"),
		Short: "DeleteSurveyProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "DeleteSurveyProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DeleteSurveyProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceNewSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &NewSurveyUserRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("NewSurveyUser"),
		Short: "NewSurveyUser RPC client",
		Long:  "Users",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "NewSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &NewSurveyUserRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyProjectRefId, cfg.FlagNamer("SurveyProjectRefId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountUserRefId, cfg.FlagNamer("SysAccountUserRefId"), "", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.SupportRoleValues, cfg.FlagNamer("SupportRoleValues"), "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.UserNeedValues, cfg.FlagNamer("UserNeedValues"), "")

	return cmd
}

func _SurveyServiceGetSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetSurveyUser"),
		Short: "GetSurveyUser RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GetSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceListSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &ListRequest{
		IdRequest: &IdRequest{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ListSurveyUser"),
		Short: "ListSurveyUser RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "ListSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &ListRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyUserId, cfg.FlagNamer("IdRequest SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountProjectId, cfg.FlagNamer("IdRequest SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyProjectId, cfg.FlagNamer("IdRequest SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountAccountId, cfg.FlagNamer("IdRequest SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.DiscoProjectId, cfg.FlagNamer("IdRequest DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountOrgId, cfg.FlagNamer("IdRequest SysAccountOrgId"), "", "")
	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")

	return cmd
}

func _SurveyServiceUpdateSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &UpdateSurveyUserRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("UpdateSurveyUser"),
		Short: "UpdateSurveyUser RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "UpdateSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &UpdateSurveyUserRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UpdateSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.SupportRoleValues, cfg.FlagNamer("SupportRoleValues"), "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.UserNeedValues, cfg.FlagNamer("UserNeedValues"), "")

	return cmd
}

func _SurveyServiceDeleteSurveyUserCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("DeleteSurveyUser"),
		Short: "DeleteSurveyUser RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "DeleteSurveyUser"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DeleteSurveyUser(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceNewDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &NewDiscoProjectRequest{
		ActionTime: &timestamp.Timestamp{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("NewDiscoProject"),
		Short: "NewDiscoProject RPC client",
		Long:  "DiscoProjects",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "NewDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &NewDiscoProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.NewDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SysAccountProjectRefId, cfg.FlagNamer("SysAccountProjectRefId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgRefId, cfg.FlagNamer("SysAccountOrgRefId"), "", "")
	cmd.PersistentFlags().StringVar(&req.Goal, cfg.FlagNamer("Goal"), "", "")
	cmd.PersistentFlags().Uint64Var(&req.AlreadyPledged, cfg.FlagNamer("AlreadyPledged"), 0, "")
	cmd.PersistentFlags().Int64Var(&req.ActionTime.Seconds, cfg.FlagNamer("ActionTime Seconds"), 0, "Represents seconds of UTC time since Unix epoch\n 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n 9999-12-31T23:59:59Z inclusive.")
	cmd.PersistentFlags().Int32Var(&req.ActionTime.Nanos, cfg.FlagNamer("ActionTime Nanos"), 0, "Non-negative fractions of a second at nanosecond resolution. Negative\n second values with fractions must still have non-negative nanos values\n that count forward in time. Must be from 0 to 999,999,999\n inclusive.")
	cmd.PersistentFlags().StringVar(&req.ActionLocation, cfg.FlagNamer("ActionLocation"), "", "")
	cmd.PersistentFlags().Uint64Var(&req.MinPioneers, cfg.FlagNamer("MinPioneers"), 0, "")
	cmd.PersistentFlags().Uint64Var(&req.MinRebelsMedia, cfg.FlagNamer("MinRebelsMedia"), 0, "")
	cmd.PersistentFlags().Uint64Var(&req.MinRebelsToWin, cfg.FlagNamer("MinRebelsToWin"), 0, "")
	cmd.PersistentFlags().StringVar(&req.ActionLength, cfg.FlagNamer("ActionLength"), "", "")
	cmd.PersistentFlags().StringVar(&req.ActionType, cfg.FlagNamer("ActionType"), "", "")
	cmd.PersistentFlags().StringVar(&req.Category, cfg.FlagNamer("Category"), "", "")
	cmd.PersistentFlags().StringVar(&req.Contact, cfg.FlagNamer("Contact"), "", "")
	cmd.PersistentFlags().StringVar(&req.HistPrecedents, cfg.FlagNamer("HistPrecedents"), "", "")
	cmd.PersistentFlags().StringVar(&req.Strategy, cfg.FlagNamer("Strategy"), "", "")
	cmd.PersistentFlags().StringSliceVar(&req.VideoUrl, cfg.FlagNamer("VideoUrl"), nil, "")
	cmd.PersistentFlags().StringVar(&req.UnitOfMeasures, cfg.FlagNamer("UnitOfMeasures"), "", "")

	return cmd
}

func _SurveyServiceGetDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetDiscoProject"),
		Short: "GetDiscoProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "GetDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}

func _SurveyServiceListDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &ListRequest{
		IdRequest: &IdRequest{},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ListDiscoProject"),
		Short: "ListDiscoProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "ListDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &ListRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyUserId, cfg.FlagNamer("IdRequest SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountProjectId, cfg.FlagNamer("IdRequest SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SurveyProjectId, cfg.FlagNamer("IdRequest SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountAccountId, cfg.FlagNamer("IdRequest SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.DiscoProjectId, cfg.FlagNamer("IdRequest DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.IdRequest.SysAccountOrgId, cfg.FlagNamer("IdRequest SysAccountOrgId"), "", "")
	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")

	return cmd
}

func _SurveyServiceUpdateDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &UpdateSurveyProjectRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("UpdateDiscoProject"),
		Short: "UpdateDiscoProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "UpdateDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &UpdateSurveyProjectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UpdateDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.SupportRoleTypes, cfg.FlagNamer("SupportRoleTypes"), "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.UserNeedTypes, cfg.FlagNamer("UserNeedTypes"), "")

	return cmd
}

func _SurveyServiceDeleteDiscoProjectCommand(cfg *client.Config) *cobra.Command {
	req := &IdRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("DeleteDiscoProject"),
		Short: "DeleteDiscoProject RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "SurveyService", "DeleteDiscoProject"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewSurveyServiceClient(cc)
				v := &IdRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DeleteDiscoProject(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.SurveyUserId, cfg.FlagNamer("SurveyUserId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountProjectId, cfg.FlagNamer("SysAccountProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SurveyProjectId, cfg.FlagNamer("SurveyProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountAccountId, cfg.FlagNamer("SysAccountAccountId"), "", "")
	cmd.PersistentFlags().StringVar(&req.DiscoProjectId, cfg.FlagNamer("DiscoProjectId"), "", "")
	cmd.PersistentFlags().StringVar(&req.SysAccountOrgId, cfg.FlagNamer("SysAccountOrgId"), "", "")

	return cmd
}
