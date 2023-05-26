// @generated by protoc-gen-connect-query v0.2.3 with parameter "target=ts"
// @generated from file proto/users/v1/users.proto (package users.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { createQueryService } from "@bufbuild/connect-query";
import { MethodKind } from "@bufbuild/protobuf";
import { CreateUserRequest, CreateUserResponse, DeleteUserRequest, DeleteUserResponse, GetUserRequest, GetUserResponse, ListUsersRequest, ListUsersResponse } from "./users_pb.js";

export const typeName = "users.v1.UserService";

/**
 * @generated from rpc users.v1.UserService.GetUser
 */
export const getUser = createQueryService({
  service: {
    methods: {
      getUser: {
        name: "GetUser",
        kind: MethodKind.Unary,
        I: GetUserRequest,
        O: GetUserResponse,
      },
    },
    typeName: "users.v1.UserService",
  },
}).getUser;

/**
 * @generated from rpc users.v1.UserService.CreateUser
 */
export const createUser = createQueryService({
  service: {
    methods: {
      createUser: {
        name: "CreateUser",
        kind: MethodKind.Unary,
        I: CreateUserRequest,
        O: CreateUserResponse,
      },
    },
    typeName: "users.v1.UserService",
  },
}).createUser;

/**
 * @generated from rpc users.v1.UserService.ListUsers
 */
export const listUsers = createQueryService({
  service: {
    methods: {
      listUsers: {
        name: "ListUsers",
        kind: MethodKind.Unary,
        I: ListUsersRequest,
        O: ListUsersResponse,
      },
    },
    typeName: "users.v1.UserService",
  },
}).listUsers;

/**
 * @generated from rpc users.v1.UserService.DeleteUser
 */
export const deleteUser = createQueryService({
  service: {
    methods: {
      deleteUser: {
        name: "DeleteUser",
        kind: MethodKind.Unary,
        I: DeleteUserRequest,
        O: DeleteUserResponse,
      },
    },
    typeName: "users.v1.UserService",
  },
}).deleteUser;
