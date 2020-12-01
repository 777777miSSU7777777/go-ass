import { Model, snakeCaseMappers, } from 'objection';

export class GenreTracks extends Model {
    genreId: string;
    trackId: string;

    static get tableName() {
        return 'genre_tracks';
    }

    static get idColumn() {
        return ['genre_id', 'track_id'];
    }

    static get columnNameMappers() {
        return snakeCaseMappers();
    }
}

export default GenreTracks;
