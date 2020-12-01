import { DBModel } from "./base";

export interface UserPlaylists extends DBModel {
    userId: string;
    playlistId: string;
}