import { types, flow } from "mobx-state-tree";
import { Views } from "./Constants";
import API from "./services/API";

const ChallangeOption = types.model({
  id: types.string,
  name: types.string,
  avatar: types.string
});

const Challange = types.model({
  id: types.string,
  text: types.string,
  options: types.array(ChallangeOption)
});

const Store = types
  .model({
    currentView: types.enumeration(Object.values(Views)),
    challanges: types.array(Challange),
    currentChallangeIdx: 0,
    error: types.maybe(types.string)
  })

  .views(self => ({
    get currentChallange() {
      return self.challanges[self.currentChallangeIdx];
    }
  }))

  .actions(self => ({
    fetchChallanges: flow(function*() {
      try {
        self.challanges = yield API.FetchChallanges();
        self.presentView(Views.PIN);
      } catch (err) {
        self.error = err;
        self.presentView(Views.ERROR);
      }
    }),

    nextChallange() {
      if (self.currentChallangeIdx < self.challanges.length - 1) {
        self.currentChallangeIdx++;
      } else {
        self.presentView(Views.SCORE);
      }
    },

    presentView(view) {
      self.currentView = view;
    }
  }));

const DefaultState = {
  currentView: Views.LOADING,
  challanges: []
};

export default Store.create(DefaultState);
