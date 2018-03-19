import { TestBed, inject } from '@angular/core/testing';

import { InitAccountService } from './init-account.service';

describe('InitAccountService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [InitAccountService]
    });
  });

  it('should be created', inject([InitAccountService], (service: InitAccountService) => {
    expect(service).toBeTruthy();
  }));
});
