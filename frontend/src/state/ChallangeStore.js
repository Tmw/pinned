import { types, flow, getRoot } from "mobx-state-tree";
import { Views } from "../Constants";
import API from "../services/API";

const Author = types.model({
  id: types.string,
  name: types.string,
  avatar: types.string
});

const Challange = types
  .model({
    id: types.string,
    text: types.string,
    options: types.array(Author),
    author: Author,
    answeredAuthorId: types.maybeNull(types.string)
  })
  .actions(self => ({
    answer(authorId) {
      const { ChallangeStore, ViewStore } = getRoot(self);

      self.answeredAuthorId = authorId;

      if (!ChallangeStore.currentChallange) {
        ViewStore.presentView(Views.SCORE);
      }
    }
  }))
  .views(self => ({
    get isAnswered() {
      return self.answeredAuthorId !== null;
    }
  }));

const ChallangeStore = types
  .model({
    challanges: types.array(Challange),
    error: types.maybe(types.string)
  })

  .views(self => ({
    get currentChallangeIdx() {
      return self.challanges.filter(challange => challange.isAnswered).length;
    },

    get currentChallange() {
      return self.challanges.find(challange => !challange.isAnswered);
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
