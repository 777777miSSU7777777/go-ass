import { DBModel } from "./base";

export interface UserTokens extends DBModel {
    userId: string;
    token: string;
}