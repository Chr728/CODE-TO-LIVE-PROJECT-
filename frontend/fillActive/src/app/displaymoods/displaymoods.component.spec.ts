import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DisplaymoodsComponent } from './displaymoods.component';

describe('DisplaymoodsComponent', () => {
  let component: DisplaymoodsComponent;
  let fixture: ComponentFixture<DisplaymoodsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DisplaymoodsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DisplaymoodsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
