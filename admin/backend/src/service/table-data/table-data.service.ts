import { Artist, Genre, GenreTracks, Playlist, PlaylistTracks, TableData, Track, User, UserPlaylists, UserTokens, UserTracks } from '@model';
import { DataRoutes } from '@app/enums';
import { Injectable } from '@nestjs/common';
import ArtistService from './artist/artist.service';
import GenreTracksService from './genre-tracks/genre-tracks.service';
import GenreService from './genre/genre.service';
import PlaylistService from './playlist/playlist.service';
import TrackService from './track/track.service';
import UserPlaylistsService from './user-playlists/user-playlists.service';
import UserTokensService from './user-tokens/user-tokens.service';
import UserTracksService from './user-tracks/user-tracks.service';
import UserService from './user/user.service';
import PlaylistTracksService from './playlist-tracks/playlist-tracks.service';

@Injectable()
export class TableDataService {
    constructor(
        private readonly artistService: ArtistService,
        private readonly genreService: GenreService,
        private readonly genreTracksService: GenreTracksService,
        private readonly playlistService: PlaylistService,
        private readonly playlistTracksService: PlaylistTracksService,
        private readonly trackService: TrackService,
        private readonly userService: UserService,
        private readonly userPlaylistsService: UserPlaylistsService,
        private readonly userTokensService: UserTokensService,
        private readonly userTracksService: UserTracksService,
    ) {}

    async getTableData(dataRoute: string): Promise<TableData[]> {
        switch(dataRoute) {
            case DataRoutes.artist:
                return await this.artistService.getArtists();
            case DataRoutes.genre:
                return await this.genreService.getGenres();
            case DataRoutes.genreTracks:
                return await this.genreTracksService.getGenreTracks();
            case DataRoutes.playlist:
                return await this.playlistService.getPlaylists();
            case DataRoutes.playlistTracks:
                return await this.playlistTracksService.getPlaylistTracks();
            case DataRoutes.track:
                return await this.trackService.getTracks();
            case DataRoutes.user:
                return await this.userService.getUsers();
            case DataRoutes.userPlaylists:
                return await this.userPlaylistsService.getUserPlaylists();
            case DataRoutes.userTokens:
                return await this.userTokensService.getUserTokens();
            case DataRoutes.userTracks:
                return await this.userTracksService.getUserTracks();
            default:
                throw new Error(`There is no table for this route '${dataRoute}'`);
        }
    }

    async newTableData(dataRoute: string, data: TableData[]): Promise<TableData[]> {
        switch(dataRoute) {
            case DataRoutes.artist:
                return await this.artistService.newArtists(data as Artist[]);
            case DataRoutes.genre:
                return await this.genreService.newGenres(data as Genre[]);
            case DataRoutes.genreTracks:
                return await this.genreTracksService.newGenreTracks(data as GenreTracks[]);
            case DataRoutes.playlist:
                return await this.playlistService.newPlaylists(data as Playlist[]);
            case DataRoutes.playlistTracks:
                return await this.playlistTracksService.newPlaylistTracks(data as PlaylistTracks[]);
            case DataRoutes.track:
                return await this.trackService.newTracks(data as Track[]);
            case DataRoutes.user:
                return await this.userService.newUsers(data as User[]);
            case DataRoutes.userPlaylists:
                return await this.userPlaylistsService.newUserPlaylists(data as UserPlaylists[]);
            case DataRoutes.userTokens:
                return await this.userTokensService.newUserTokens(data as UserTokens[]);
            case DataRoutes.userTracks:
                return await this.userTracksService.newUserTracks(data as UserTracks[]);
            default:
                throw new Error(`There is no table for this route '${dataRoute}'`);
        }
    }

    async updateTableData(dataRoute: string, data: TableData[]): Promise<TableData[]> {
        switch(dataRoute) {
            case DataRoutes.artist:
                return await this.artistService.updateArtists(data as Artist[]);
            case DataRoutes.genre:
                return await this.genreService.updateGenres(data as Genre[]);
            case DataRoutes.genreTracks:
                return await this.genreTracksService.updateGenreTracks(data as GenreTracks[]);
            case DataRoutes.playlist:
                return await this.playlistService.updatePlaylists(data as Playlist[]);
            case DataRoutes.playlistTracks:
                return await this.playlistTracksService.updatePlaylistTracks(data as PlaylistTracks[]);
            case DataRoutes.track:
                return await this.trackService.updateTracks(data as Track[]);
            case DataRoutes.user:
                return await this.userService.updateUsers(data as User[]);
            case DataRoutes.userPlaylists:
                return await this.userPlaylistsService.updateUserPlaylists(data as UserPlaylists[]);
            case DataRoutes.userTokens:
                return await this.userTokensService.updateUserTokens(data as UserTokens[]);
            case DataRoutes.userTracks:
                return await this.userTracksService.updateUserTracks(data as UserTracks[]);
            default:
                throw new Error(`There is no table for this route '${dataRoute}'`);
        }
    }

    async deleteTableData(dataRoute: string, data: TableData[]): Promise<number[]> {
        switch(dataRoute) {
            case DataRoutes.artist:
                return await this.artistService.deleteArtists(data as Artist[]);
            case DataRoutes.genre:
                return await this.genreService.deleteGenres(data as Genre[]);
            case DataRoutes.genreTracks:
                return await this.genreTracksService.deleteGenreTracks(data as GenreTracks[]);
            case DataRoutes.playlist:
                return await this.playlistService.deletePlaylists(data as Playlist[]);
            case DataRoutes.playlistTracks:
                return await this.playlistTracksService.deletePlaylistTracks(data as PlaylistTracks[]);
            case DataRoutes.track:
                return await this.trackService.deleteTracks(data as Track[]);
            case DataRoutes.user:
                return await this.userService.deleteUsers(data as User[]);
            case DataRoutes.userPlaylists:
                return await this.userPlaylistsService.deleteUserPlaylists(data as UserPlaylists[]);
            case DataRoutes.userTokens:
                return await this.userTokensService.deleteUserTokens(data as UserTokens[]);
            case DataRoutes.userTracks:
                return await this.userTracksService.deleteUserTracks(data as UserTracks[]);
            default:
                throw new Error(`There is no table for this route '${dataRoute}'`);
        }
    }
}

export default TableDataService;
