import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UsabilityTestComponent } from './usability-test.component';

describe('UsabilityTestComponent', () => {
  let component: UsabilityTestComponent;
  let fixture: ComponentFixture<UsabilityTestComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UsabilityTestComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UsabilityTestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
