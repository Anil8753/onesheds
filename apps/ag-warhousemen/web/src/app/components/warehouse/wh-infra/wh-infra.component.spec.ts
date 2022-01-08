import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhInfraComponent } from './wh-infra.component';

describe('WhInfraComponent', () => {
  let component: WhInfraComponent;
  let fixture: ComponentFixture<WhInfraComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WhInfraComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WhInfraComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
