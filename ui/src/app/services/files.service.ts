import { HttpClient, HttpEventType } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class FilesService {
  constructor(private http: HttpClient) {}
  list(path: string) {
    return this.http.get<FileInfo[]>('api/file', { params: { path } });
  }
  move(old_path: string, new_path: string) {
    return this.http.post('api/file/move', { old_path, new_path });
  }
  mkdir(path: string) {
    return this.http.post('api/file/mkdir', { path });
  }
  delete(path: string) {
    return this.http.delete('api/file', { params: { path } });
  }
  copy(old_path: string, new_path: string) {
    return this.http.post('api/file/copy', { old_path, new_path });
  }
  upload(path: string, file: File, callback?: (progress: number) => void) {
    const formdata = new FormData();
    formdata.append('path', path);
    formdata.append('file', file);
    return new Promise((resolve, reject) => {
      this.http
        .post('api/file', formdata, { reportProgress: true, observe: 'events' })
        .subscribe(
          (event) => {
            console.log(event);
            if (event.type === HttpEventType.UploadProgress && event.total) {
              if (callback) {
                callback((event.loaded / event.total) * 100);
              }
            } else if (event.type === HttpEventType.Response) {
              resolve(null);
            }
          },
          (err) => reject(err)
        );
    });
  }

  baseDir(path: string) {
    if (path[path.length - 1] === '/') {
      path = path.slice(0, -1);
    }
    const index = path.lastIndexOf('/');
    return index !== -1 ? path.slice(0, index) || '/' : '';
  }
  baseName(path: string) {
    if (path[path.length - 1] === '/') {
      path = path.slice(0, -1);
    }
    const index = path.lastIndexOf('/');
    return index !== -1 ? path.slice(index + 1) : path;
  }
}

export interface FileInfo {
  name: string;
  fullname: string;
  size: number;
  mode: number;
  mod_time: string;
  is_dir: boolean;
  ext: string;
}
