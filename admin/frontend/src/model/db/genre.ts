import { DBModel } from "./base";

export interface Genre extends DBModel  {
    genreId: string;
    genreTitle: string;
}