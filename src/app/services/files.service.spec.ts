import { TestBed } from '@angular/core/testing';
import {
  HttpClientTestingModule,
  HttpTestingController,
} from '@angular/common/http/testing';
import { FilesService } from './files.service';

describe('FilesService', () => {
  let service: FilesService;
  let backend: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(FilesService);
    backend = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
  it('basename', () => {
    expect(service.baseName('/abc/def')).toEqual('def');
  });
  it('basedir', () => {
    expect(service.baseDir('abc/def')).toEqual('abc');
    expect(service.baseDir('/abc/def')).toEqual('/abc');
    expect(service.baseDir('/abc')).toEqual('/');
    expect(service.baseDir('/abc/def/ghi')).toEqual('/abc/def');
    expect(service.baseDir('/abc/def/ghi/')).toEqual('/abc/def');
  });
});
