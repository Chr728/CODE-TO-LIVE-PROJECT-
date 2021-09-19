import { Component, OnInit } from '@angular/core';
import {ModalDismissReasons, NgbModal} from '@ng-bootstrap/ng-bootstrap';
import { HttpClient } from "@angular/common/http";
import { NgForm} from "@angular/forms";

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
  selector: 'app-mood',
  templateUrl: './mood.component.html',
  styleUrls: ['./mood.component.css']
})
export class MoodComponent implements OnInit {
  resp: Response | undefined;
  responsedata: Data[] | undefined;
  closeResult: string | undefined;

  constructor(
    private httpClient: HttpClient,
    private modalService: NgbModal
  ) { }

  ngOnInit(): void {
  }


  open(content: any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
      console.log(this.closeResult)
    });
  }


  open2(content: any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
      console.log(this.closeResult)
    });
  }


  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {

      return 'by clicking on a backdrop';
    } else {

      return `with: ${reason}`;
    }
  }

  newWrite: W | undefined;
  message: string | undefined;
  onSubmit(f: NgForm) {
    const url = 'http://localhost:8008/playlists';
    this.httpClient.post<any>(url, f.value)
      .subscribe((result) => {
        this.resp = result;
        if (this.resp?.success == false){
          this.newWrite = result;
          this.message = this.newWrite?.message;
          console.log(this.message)
        }else{
          this.responsedata = this.resp?.data;
          console.log(this.responsedata);
        }
        this.ngOnInit(); //reload the table
      });
    this.modalService.dismissAll(); //dismiss the modal
  }

}
