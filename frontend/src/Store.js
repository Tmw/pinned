import { decorate, observable, action, computed, runInAction } from "mobx";
import { Views } from "./Constants";
import API from "./services/API";

class Store {
  /* Own properties */
  currentView = Views.LOADING;
  error = "";
  challanges = [];
  currentChallangeIdx = 0;

  /* Computed Properties */
  get currentChallange() {
    return this.challanges[this.currentChallangeIdx];
  }

  nextChallange() {
    if (this.currentChallangeIdx < this.challanges.length - 1) {
      this.currentChallangeIdx++;
    } else {
      this.showNextView();
    }
  }

  /* Actions */
  fetchChallanges() {
    this.currentView = Views.LOADING;

    API.FetchChallanges()
      .then(c => {
        runInAction(() => {
          this.challanges = c;
          this.currentView = Views.PIN;
        });
      })
      .catch(err => {
        runInAction(() => {
          this.currentView = Views.ERROR;
          this.error = err;
        });
      });
  }

  showNextView() {
    switch (this.currentView) {
      case Views.LOADING:
        this.currentView = Views.PIN;
        break;

      case Views.PIN:
        this.currentView = Views.SCORE;
        break;

      default:
      case Views.SCORE:
        this.currentView = Views.LOADING;
        break;
    }
  }
}

decorate(Store, {
  currentView: observable,
  challanges: observable,
  currentChallangeIdx: observable,

  currentChallange: computed,

  showNextView: action,
  fetchChallanges: action
});

export default new Store();
