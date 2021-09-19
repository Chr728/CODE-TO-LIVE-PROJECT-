import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {HomeComponent} from "./home/home.component";
import {MoodComponent} from "./mood/mood.component";
import {DisplaymoodsComponent} from "./displaymoods/displaymoods.component";
import {AboutComponent} from "./about/about.component";


const routes: Routes = [
  {path: "", component: HomeComponent},
  {path: "mood", component: MoodComponent},
  {path: "displaymood", component: DisplaymoodsComponent},
  {path: "about", component:AboutComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
