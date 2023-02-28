import { TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpTestingController } from '@angular/common/http/testing'; 
import { MixService } from './mix.service';

describe('MixService', () => {
  let service: MixService;
  let httpClient: HttpClientModule;
  let httpTestingController: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientModule, HttpClientTestingModule],
      providers: [MixService]
    })
    TestBed.configureTestingModule({});
    service = TestBed.inject(MixService);
    httpClient = TestBed.inject(HttpClientModule);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('getList function makes post request', () => {
    const results: number[] = [1, 2, 3, 4];

    service.getList(results);

    const req = httpTestingController.expectOne('http://localhost:8080/results');
    expect(req.request.method).toEqual('POST');
    expect(req.request.body).toEqual(results);
  });

});
