import { Component, OnInit } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {NgbModal} from "@ng-bootstrap/ng-bootstrap";
import {NgForm} from "@angular/forms";





export class ExternalUrls {
  constructor(
    public spotify: string,
  ) {
  }
}
export class Data {
  constructor(
    public description: string,
    public external_urls: ExternalUrls,
    public id: string,
    public name: string,
    public event: string,
    public image_url: string,
    public images: Images[],
  ) {
  }
}
export class Response {
  constructor(
    public success: boolean,
    public data: Data[],
  ) {
  }
}
export class Images {
  constructor(
    public height: string,
    public Url: string,
    public width: string,
  ) {
  }
}


export class W {
  constructor(
    public success: boolean,
    public message: string,
  ) {
  }
}




@Component({
  selector: 'app-displaymoods',
  templateUrl: './displaymoods.component.html',
  styleUrls: ['./displaymoods.component.css']
})
export class DisplaymoodsComponent implements OnInit {
  url = 'http://localhost:8008/suggestplayist'
  resp: Response | undefined;
  responsedata: Data[] | undefined;
  closeResult: string | undefined;

  constructor(
    private httpClient: HttpClient,
    private modalService: NgbModal
  ) { }

  ngOnInit(): void {
    this.getResponse()
  }



  newWrite: W | undefined;
  message: string | undefined;
  getResponse(){
    this.httpClient.get<any>(this.url).subscribe(
      (result) => {
        this.resp = result;
          this.responsedata = this.resp?.data;
          console.log(this.responsedata);
          alert(this.responsedata);
      }
    )
  }
}


