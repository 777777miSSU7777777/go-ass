import { Model, snakeCaseMappers } from 'objection';

export class Genre extends Model {
    genreId: string;
    genreTitle: string;
    
    static get tableName() {
        return 'genres';
    }

    static get idColumn() {
        return 'genre_id';
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default Genre;