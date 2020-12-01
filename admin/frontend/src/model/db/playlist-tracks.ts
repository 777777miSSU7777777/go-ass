import { DBModel } from "./base";

export interface PlaylistTracks extends DBModel  {
    playlistId: string;
    trackId: string;
}