import type { IReport } from "@/models/report";

export interface IUser {
  username: string;
  name: string;
  email: string;
}

export interface GenericResponse {
  status: string;
  message: string;
}

export interface ILoginInput {
  username: string;
  password: string;
}

export interface ISignUpInput {
  username: string;
  email: string;
  password: string;
  name: string;
}

export interface ILoginResponse {
  status: string;
  message: string;
  access_token: string;
}

export interface ISignUpResponse {
  status: string;
  message: string;
}

export interface IUserResponse {
  status: string;
  data: IUser;
}

export interface IReportResponse {
  status: string;
  data: IReport[];
  total: Number;
}
