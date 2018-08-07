import { types, flow, getRoot } from "mobx-state-tree";
import { Views } from "../Constants";
import API from "../services/API";

const Author = types.model({
  id: types.string,
  name: types.string,
  avatar: types.string
});

const Challange = types.model({
  id: types.string,
  text: types.string,
  options: types.array(Author),
  author: Author,
  answeredAuthorId: types.maybeNull(types.string)
});

const ChallangeStore = types
  .model({
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
        getRoot(self).ViewStore.presentView(Views.PIN);
      } catch (err) {
        self.error = err;
        getRoot(self).ViewStore.presentView(Views.ERROR);
      }
    }),

    nextChallange() {
      if (self.currentChallangeIdx < self.challanges.length - 1) {
        self.currentChallangeIdx++;
      } else {
        getRoot(self).ViewStore.presentView(Views.SCORE);
      }
    }
  }));

const Defaults = {
  challanges: []
};

const CreateChallangeStore = overrides =>
  ChallangeStore.create(Object.assign(Defaults, overrides));

export { ChallangeStore, CreateChallangeStore };
