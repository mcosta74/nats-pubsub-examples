import { TestBed } from '@angular/core/testing';

import { NatsService } from './nats.service';

describe('NatsService', () => {
  let service: NatsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NatsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
