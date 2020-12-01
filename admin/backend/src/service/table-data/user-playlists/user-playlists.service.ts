import { Injectable } from '@nestjs/common';
import { UserPlaylists } from "@model";

@Injectable()
export class UserPlaylistsService {
  async getUserPlaylists(): Promise<UserPlaylists[]> {
    return await UserPlaylists.query().select();
  }

  async newUserPlaylists(userPlaylists: UserPlaylists[]): Promise<UserPlaylists[]> {
    return await UserPlaylists.query().insertAndFetch(userPlaylists);
  }  

  async updateUserPlaylists(userPlaylists: UserPlaylists[]): Promise<UserPlaylists[]> {
    return Promise.all(userPlaylists.map((userPlaylist: UserPlaylists) => {
      return UserPlaylists.query().updateAndFetchById(userPlaylist.userId, userPlaylist);
    }));
  }

  async deleteUserPlaylists(userPlaylists: UserPlaylists[]): Promise<number[]> {
    return Promise.all(userPlaylists.map((userPlaylist: UserPlaylists) => {
      return UserPlaylists.query().deleteById([userPlaylist.playlistId, userPlaylist.playlistId]);
    }))
  }
}

export default UserPlaylistsService;