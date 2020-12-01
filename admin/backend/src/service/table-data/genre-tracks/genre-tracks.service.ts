import { Injectable } from '@nestjs/common';
import { GenreTracks } from "@model";

@Injectable()
export class GenreTracksService {
  async getGenreTracks(): Promise<GenreTracks[]> {
    return await GenreTracks.query().select();
  }

  async newGenreTracks(genreTracks: GenreTracks[]): Promise<GenreTracks[]> {
    return await GenreTracks.query().insertAndFetch(genreTracks);
  }

  async updateGenreTracks(genreTracks: GenreTracks[]): Promise<GenreTracks[]> {
    return Promise.all(genreTracks.map((genreTrack: GenreTracks) => {
      return GenreTracks.query().updateAndFetchById(genreTrack.genreId, genreTrack);
    }));
  }

  async deleteGenreTracks(genreTracks: GenreTracks[]): Promise<number[]> {
    return Promise.all(genreTracks.map((genreTrack: GenreTracks) => {
      return GenreTracks.query().deleteById([genreTrack.genreId, genreTrack.trackId]);
    }))
  }
}

export default GenreTracksService;