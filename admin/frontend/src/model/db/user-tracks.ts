import { DBModel } from "./base";

export interface UserTracks extends DBModel {
    userId: string;
    trackId: string;
}