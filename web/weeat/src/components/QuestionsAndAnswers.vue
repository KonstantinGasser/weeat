<template>
  <div class="content">
    <div class="headline text-center">
      <h1>
        <a href="/"> <span>we</span>Eat </a>
      </h1>
    </div>
    <div class="random-stats">
      <div class="stat foods">
        <span class="stat-name">#foods</span>
        <span class="stat-figure">1024</span>
      </div>
      <div class="stat foods">
        <span class="stat-name">#recipes</span>
        <span class="stat-figure">256</span>
      </div>
    </div>
    <section>
      <div class="space"></div>
      <h2>
        What the heck is <span class="inline-name"><span>we</span>Eat</span>
      </h2>
      <div class="info-card" @click="setActive(0)">
        <h3>Idea behind the App..</h3>
        <small>
          <i v-if="activeBox !== 0" class="bi bi-caret-right-fill">show more</i>
          <i v-if="activeBox === 0" class="bi bi-caret-down-fill">show less</i>
        </small>
        <p :class="{ expand: activeBox === 0 }"></p>
      </div>
      <div class="info-card" @click="setActive(1)">
        <h3>Free for all ✨</h3>
        <small>
          <i v-if="activeBox !== 1" class="bi bi-caret-right-fill">show more</i>
          <i v-if="activeBox === 1" class="bi bi-caret-down-fill">show less</i>
        </small>
      </div>

      <div class="space"></div>
      <h2>What do we store about you?</h2>
      <div class="info-card" @click="setActive(10)">
        <h3>without cookies 🚫🍪</h3>
        <small>
          <i v-if="activeBox !== 10" class="bi bi-caret-right-fill"
            >show more</i
          >
          <i v-if="activeBox === 10" class="bi bi-caret-down-fill">show less</i>
        </small>
        <p :class="{ expand: activeBox === 10 }">
          <span class="inline-name"><span>we</span>Eat</span> only stores
          information about you if you accept the requested cookie. If you chose
          not to then no cookie will be stored neither on the client device nor
          a session gets created on the server side. Further, if no cookie is
          present the server will also
          <span class="text-underline text-emph">not</span> track any of your
          personal information such as your IP address, your device information
          (which browser, operating system, resolution) - there is
          <span class="text-underline text-emph">not</span> digital-fingerprint
          created about you. <br /><br />
          <span class="text-emph">Why?</span>
          <br />
          👉 because it is not required in order to run the application
          smoothly. And the last thing I care about is your IP ~ sorry 😅
        </p>
      </div>
      <div class="info-card" @click="setActive(11)">
        <h3>with cookies 🍪</h3>
        <small>
          <i v-if="activeBox !== 11" class="bi bi-caret-right-fill"
            >show more</i
          >
          <i v-if="activeBox === 11" class="bi bi-caret-down-fill">show less</i>
        </small>
        <p :class="{ expand: activeBox === 11 }">
          If you accept the cookie however,
          <span class="inline-name"><span>we</span>Eat</span> will store limited
          information about you. A cookie with a randomly created unique ID will
          be assigned to you (aka the cookie) which in turn will enable you to
          track your intakes for the current day. On the server side any food or
          recipe you add to your daily intake will be mapped to the cookie and
          stored in the database. <br /><br />
          <span class="text-emph">Important!?</span>
          <br />
          👉 The cookie does also
          <span class="text-underline text-emph">not</span> include any personal
          information (IP, device, OS, etc.) about you. It is only a random
          "username" you get assigned for this day. <br /><br />
          👉 Every day at <span class="text-emph">00:00 o'clock</span> the
          cookie gets invalidated and all corresponding data stored in the
          database will be delete.
          <small
            >(This is also the reason why you get asked to accept or reject the
            cookie every day...)</small
          >
        </p>
      </div>
      <div class="space"></div>
      <h2>How content is verified?</h2>
      <div class="info-card" @click="setActive(20)">
        <h3>Everyone can create content 🗣</h3>
        <small>
          <i v-if="activeBox !== 20" class="bi bi-caret-right-fill"
            >show more</i
          >
          <i v-if="activeBox === 20" class="bi bi-caret-down-fill">show less</i>
        </small>
        <p :class="{ expand: activeBox === 20 }">
          The overall goal of
          <span class="inline-name"><span>we</span>Eat</span> is it to create an
          application which enables you to quickly get an idea about a food's
          nutrition values or to get inspired by recipes others have shared.
          Sure there are many apps that can do that if you like the pay-wall 💸
          they all have.. <br /><br />
          <span class="text-emph">Now, where is the problem?</span>
          <br />
          If you have a network where literally anyone can share food nutrition
          values with no central authority you eventually will have data which
          not only seems wrong but is indeed wrong. Somehow
          <span class="inline-name"><span>we</span>Eat</span>
          needs to address this issue..

          <br /><br />
          <span class="text-emph">Verification process</span>
          <br />
          <span class="inline-name"><span>we</span>Eat</span> mitigates this
          issue by having a defined verification process. After a food item has
          been created the process gets triggered and the newly added item runs
          through the verification process. Randomly, users opening the
          application will be asked to verify the newly added item. Unless
          <span class="text-emph">51%</span>
          of the requested users acknowledge that the new item is in-fact
          correct the item will be marked as verified ✅
        </p>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "PrivacyNotice",
  data() {
    return {
      activeBox: -1, // start at -1 since index starts at zero (0) -> init is inactive
    };
  },
  methods: {
    setActive(id) {
      if (this.activeBox === id) {
        this.activeBox = -1; // set inactive
        return;
      }
      this.activeBox = id;
    },
  },
};
</script>

<style scoped>
.content {
  padding: 15px 0px;
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}

section {
  width: 100vw;
  height: 100%;
  margin-bottom: 75px;
  overflow: scroll;
  margin-bottom: 0px;
  display: flex;
  flex-direction: column;
  gap: 15px;
  padding: 15px 10px;
}

section .info-card {
  box-shadow: 0 0 10px 2px rgb(0 0 0 / 10%);
  background: var(--box-bg);
  border-radius: 14px;
  padding: 15px;
}

section .info-card p {
  opacity: 0;
  height: 0px;
  visibility: hidden;
  transition: height 100ms linear;
  transition: opacity 50ms linear;
}

section .info-card > p.expand {
  height: max-content;
  visibility: visible;
  opacity: 1;
  transition: opacity 300ms linear;
}

section .space {
  margin: 10px 0;
}


.random-stats {
  display: flex;
  justify-content: space-evenly;
  padding: 10px 10px;
}

.random-stats .stat {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: var(--box-bg);
  padding: 5px 15px;
  border-radius: 14px;
  box-shadow: 0 0 6px 2px rgba(0, 0, 0, 0.1);
}

.stat .stat-name,
.stat-figure {
  width: max-content;
  font-weight: bold;
  color: #000000;
}
.stat .stat-name {
  font-size: 20px;
}
.stat .stat-figure {
  font-size: 24px;
}
</style>