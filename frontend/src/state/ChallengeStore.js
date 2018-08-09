import { types, flow, getRoot } from "mobx-state-tree";
import { Views } from "../Constants";
import API from "../services/API";

const Author = types.model({
  id: types.string,
  name: types.string,
  avatar: types.string
});

const Challenge = types
  .model({
    id: types.string,
    text: types.string,
    options: types.array(Author),
    author: Author,
    answeredAuthorId: types.maybeNull(types.string)
  })
  .actions(self => ({
    answer(authorId) {
      const { ChallengeStore, ViewStore } = getRoot(self);

      self.answeredAuthorId = authorId;

      if (!ChallengeStore.currentChallenge) {
        ViewStore.presentView(Views.SCORE);
      }
    }
  }))
  .views(self => ({
    get isAnswered() {
      return self.answeredAuthorId !== null;
    },

    get isCorrect() {
      return self.answeredAuthorId === self.author.id;
    }
  }));

const ChallengeStore = types
  .model({
    challenges: types.array(Challenge)
  })

  .views(self => ({
    get currentChallengeIndex() {
      return self.challenges.filter(challenge => challenge.isAnswered).length;
    },

    get currentChallenge() {
      return self.challenges.find(challenge => !challenge.isAnswered);
    }
  }))

  .actions(self => ({
    fetchChallenges: flow(function*() {
      try {
        self.challenges = yield API.FetchChallenges();
        getRoot(self).ViewStore.presentView(Views.PIN);
      } catch (err) {
        getRoot(self).ViewStore.presentView(Views.ERROR, err.message);
      }
    })
  }));

const Defaults = {
  challenges: []
};

const CreateChallengeStore = overrides =>
  ChallengeStore.create(Object.assign(Defaults, overrides));

export { ChallengeStore, CreateChallengeStore };
