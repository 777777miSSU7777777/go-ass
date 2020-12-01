import { Injectable } from '@nestjs/common';
import { Artist } from "@model";
import * as uuid from 'uuid';

@Injectable()
export class ArtistService {
  async getArtists(): Promise<Artist[]> {
    return await Artist.query().select();
  }

  async newArtists(artists: Artist[]): Promise<Artist[]> {
      return Artist.query().insertAndFetch(artists.map((artist: Artist) => {
        return { 
            ...artist,
            artistId: artist.artistId || uuid.v4(),
        }
    }));
  }

  async updateArtists(artists: Artist[]): Promise<Artist[]> {
    return Promise.all(artists.map((artist: Artist) => {
      return Artist.query().updateAndFetchById(artist.artistId, artist)
    }));
  }

  async deleteArtists(artists: Artist[]): Promise<number[]> {
    return Promise.all(artists.map((artist: Artist) => {
      return Artist.query().deleteById(artist.artistId);
    }))
  }
}

export default ArtistService;
