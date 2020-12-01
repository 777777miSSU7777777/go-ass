import { Injectable } from '@nestjs/common';
import { UserTokens } from "@model";

@Injectable()
export class UserTokensService {
  async getUserTokens(): Promise<UserTokens[]> {
    return await UserTokens.query().select();
  }

  async newUserTokens(userTokens: UserTokens[]): Promise<UserTokens[]> {
    return await UserTokens.query().insertAndFetch(userTokens);
  }  

  async updateUserTokens(userTokens: UserTokens[]): Promise<UserTokens[]> {
    return Promise.all(userTokens.map((userToken: UserTokens) => {
      return UserTokens.query().updateAndFetchById(userToken.userId, userToken);
    }));
  }

  async deleteUserTokens(userTokens: UserTokens[]): Promise<number[]> {
    return Promise.all(userTokens.map((userTokens: UserTokens) => {
      return UserTokens.query().deleteById([userTokens.userId, userTokens.token]);
    }))
  }
}

export default UserTokensService;