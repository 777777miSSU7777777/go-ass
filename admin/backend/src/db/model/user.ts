import { Model, snakeCaseMappers } from 'objection';

export class User extends Model {
    userId: string;
    role: string;
    email: string;
    username: string;
    password: string;
    
    static get tableName() {
        return 'users';
    }

    static get idColumn() {
        return 'user_id';
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default User;