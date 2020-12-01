import { DBModel } from "./base";

export interface Track extends DBModel {
    trackId: string;
    trackTitle: string;
    artistId: string;
    genreId: string;
    uploadedById: string;
}