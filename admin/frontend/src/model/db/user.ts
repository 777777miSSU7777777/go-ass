import { DBModel } from "./base";

export interface User extends DBModel {
    userId: string;
    role: string;
    email: string;
    username: string;
    password: string;
}