import { DBModel } from "./base";

export interface Playlist extends DBModel  {
    playlistId: string;
    playlistTitle: string;
    createdById: string;
}