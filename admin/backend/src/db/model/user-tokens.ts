import { Model, snakeCaseMappers } from 'objection';

export class UserTokens extends Model {
    userId: string;
    token: string;
    
    static get tableName() {
        return 'user_tokens';
    }

    static get idColumn() {
        return ['user_id', 'token'];
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default UserTokens;