import { DBModel } from "./base";

export interface Artist extends DBModel {
    artistId: string;
    artistName: string;
}