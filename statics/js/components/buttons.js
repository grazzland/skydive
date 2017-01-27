/* jshint multistr: true */
Vue.component('button-state', {

  props: {

    value: {
      type: Boolean,
      required: true,
    },

    enabledText: {
      type: String,
      required: true,
    },

    disabledText: {
      type: String,
      required: true,
    }

  },

  template: '\
    <button type="button" class="btn btn-default"\
            :class="{\'active\': value}"\
            @click="change">\
      <span v-if="value">{{enabledText}}</span>\
      <span v-else>{{disabledText}}</span>\
    </button>\
  ',

  methods: {

    change: function() {
      this.$emit('input', !this.value);
    }

  }

});

Vue.component('button-dropdown', {

  props: {

    text: {
      type: String,
    },

    bClass: {
      type: String,
    },

    dropup: {
      type: Boolean,
      default: false,
    },

    autoClose: {
      type: Boolean,
      default: true,
    },

  },

  template: '\
    <div class="btn-group" :class="{\'open\': open, \'dropup\': dropup}">\
      <button class="btn btn-default dropdown-toggle"\
              :class="bClass"\
              @click="toggle"\
              aria-haspopup="true"\
              :aria-expanded="open">\
          <slot name="button-text">\
            {{text}}\
          </slot>\
      </button>\
      <ul class="dropdown-menu" v-if="open" @click="itemSelected">\
        <slot></slot>\
      </ul>\
    </div>\
  ',

  data: function() {
    return {
      open: false,
    };
  },

  mounted: function() {
    var self = this;
    // close the popup if we click elsewhere
    $(document).on('click', function(event) {
      if (!$(event.target).parents('.btn-group').length) {
        self.open = false;
      }
    });
  },

  methods: {

    toggle: function() {
      this.open = !this.open;
    },

    itemSelected: function() {
      if (this.autoClose) {
        this.toggle();
      }
    },

  },

});
