import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DonorLoginComponent } from './donor-login.component';

describe('DonorLoginComponent', () => {
  let component: DonorLoginComponent;
  let fixture: ComponentFixture<DonorLoginComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DonorLoginComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DonorLoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
