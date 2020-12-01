import { DBModel } from "./base";

export interface GenreTracks extends DBModel   {
    genreId: string;
    trackId: string;
}