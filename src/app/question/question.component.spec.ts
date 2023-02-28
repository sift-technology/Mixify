import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed, waitForAsync, fakeAsync, tick, } from '@angular/core/testing';
import { RouterTestingModule } from "@angular/router/testing";
import { Router } from "@angular/router";
import { AppModule } from '../app.module';
import { Location } from '@angular/common';

import { QuestionComponent } from './question.component';
import { ResultsComponent } from '../results/results.component';

describe('QuestionComponent', () => {
  let router: Router;
  let location: Location;
  let component: QuestionComponent;
  let fixture: ComponentFixture<QuestionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ QuestionComponent ],
      imports: [HttpClientModule, AppModule, RouterTestingModule.withRoutes([
        { path: '', component: QuestionComponent },
        { path: 'results', component: ResultsComponent }
      ])],
    })
    .compileComponents();

    router = TestBed.inject(Router);
    location = TestBed.inject(Location);

    fixture = TestBed.createComponent(QuestionComponent);
    router.initialNavigation();
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });

  it('title in h2 tag is "Questions: " ', waitForAsync(() => {
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h2').textContent).toContain('Questions: ');
  }));

  it('click "Submit" calls onSubmit function', fakeAsync(() => {
    spyOn(component, 'onSubmit');
    let button = fixture.debugElement.nativeElement.querySelector('#submit');
    button.click();
    tick();
    expect(component.onSubmit).toHaveBeenCalled();
  }));

});
