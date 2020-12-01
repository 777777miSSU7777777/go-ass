import { Injectable } from '@nestjs/common';
import { Genre } from "@model";
import * as uuid from 'uuid';

@Injectable()
export class GenreService {
  async getGenres(): Promise<Genre[]> {
    return await Genre.query().select();
  }

  async newGenres(genres: Genre[]): Promise<Genre[]> {
    return await Genre.query().insertAndFetch(genres.map((genre: Genre) => {
      return {
        ...genre,
        genreId: genre.genreId || uuid.v4(),
      };
    }));
  }

  async updateGenres(genres: Genre[]): Promise<Genre[]> {
    return Promise.all(genres.map((genre: Genre) => {
      return Genre.query().updateAndFetchById(genre.genreId, genre);
    }));
  }
  
  async deleteGenres(genres: Genre[]): Promise<number[]> {
    return Promise.all(genres.map((genre: Genre) => {
      return Genre.query().deleteById(genre.genreId);
    }))
  }
}

export default GenreService;