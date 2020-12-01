import { Injectable } from '@nestjs/common';
import { Playlist } from "@model";
import * as uuid from 'uuid';

@Injectable()
export class PlaylistService {
  async getPlaylists(): Promise<Playlist[]> {
    return await Playlist.query().select();
  }
  
  async newPlaylists(playlists: Playlist[]): Promise<Playlist[]> {
    return await Playlist.query().insertAndFetch(playlists.map((playlist: Playlist) => {
      return {
        ...playlist,
        playlistId: playlist.playlistId || uuid.v4(),
      };
    }));
  }

  async updatePlaylists(playlists: Playlist[]): Promise<Playlist[]> {
    return Promise.all(playlists.map((playlist: Playlist) => {
      return Playlist.query().updateAndFetchById(playlist.playlistId, playlist);
    }));
  }

  async deletePlaylists(playlists: Playlist[]): Promise<number[]> {
    return Promise.all(playlists.map((playlist: Playlist) => {
      return Playlist.query().deleteById(playlist.playlistId);
    }))
  }
}

export default PlaylistService;