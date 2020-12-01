import { Injectable } from '@nestjs/common';
import { UserTracks } from "@model";

@Injectable()
export class UserTracksService {
  async getUserTracks(): Promise<UserTracks[]> {
    return await UserTracks.query().select();
  }

  async newUserTracks(userTracks: UserTracks[]): Promise<UserTracks[]> {
    return UserTracks.query().insertAndFetch(userTracks);
  }  

  async updateUserTracks(userTracks: UserTracks[]): Promise<UserTracks[]> {
    return Promise.all(userTracks.map((userTrack: UserTracks) => {
      return UserTracks.query().updateAndFetchById(userTrack.userId, userTrack);
    }));
  }

  async deleteUserTracks(userTracks: UserTracks[]): Promise<number[]> {
    return Promise.all(userTracks.map((userTrack: UserTracks) => {
      return UserTracks.query().deleteById([userTrack.userId, userTrack.trackId]);
    }))
  }
}

export default UserTracksService;