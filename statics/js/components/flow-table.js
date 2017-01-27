/* jshint multistr: true */

var TableHeader = {
  
  props: {
    fields: {
      type: Array,
      required: true,
    },

    sortBy: {
      type: Array,
      required: true,
    },

    sortOrder: {
      type: Number,
      required: true,
    },

  },

  template: '\
    <thead>\
      <tr>\
        <th v-for="field in fields"\
            v-if="field.show"\
            @click="sort(field.name)"\
            >\
          {{field.label}}\
          <i v-if="field.name == sortBy"\
             style="float: right"\
             class="fa"\
             :class="{\'fa-chevron-down\': sortOrder == 1,\
                      \'fa-chevron-up\': sortOrder == -1}"\
             aria-hidden="true"></i>\
        </th>\
      </tr>\
    </thead>\
  ',

  methods: {

    sort: function(name) {
      this.$emit('sort', name);
    },

  }

};


Vue.component('flow-table', {
  
  props: {
    value: {
      type: String,
      required: true
    },
  },

  template: '\
    <div v-if="!queryError" class="flow-table">\
      <div class="flow-table-wrapper">\
        <table class="table table-condensed table-bordered">\
          <table-header :fields="fields"\
                        :sortOrder="sortOrder"\
                        :sortBy="sortBy"\
                        @sort="sort"></table-header>\
          <tbody>\
            <template v-for="obj in sortedResults">\
              <tr class="flow-row"\
                  :class="{\'flow-detail\': showDetail[obj[showDetailField]]}"\
                  @click="showFlowDetail(obj)"\
                  @mouseenter="highlightNodes(obj, true)"\
                  @mouseleave="highlightNodes(obj, false)">\
                <td v-for="field in fields" v-if="field.show">\
                  {{fieldValue(obj, field.name)}}\
                </td>\
              </tr>\
              <tr class="flow-detail-row"\
                  v-if="showDetail[obj[showDetailField]]"\
                  @mouseenter="highlightNodes(obj, true)"\
                  @mouseleave="highlightNodes(obj, false)">\
                <td :colspan="fields.length">\
                  <object-detail :object="obj"></object-detail>\
                </td>\
              </tr>\
            </template>\
          </tbody>\
        </table>\
      </div>\
      <div class="actions">\
        <button-dropdown b-class="btn-xs" :dropup="true" :auto-close="false">\
          <span slot="button-text">\
            <i class="fa fa-cog" aria-hidden="true"></i>\
          </span>\
          <li v-for="field in fields">\
            <a href="#" @click="field.show = !field.show">\
              <i class="fa fa-check text-success right"\
                 aria-hidden="true" v-show="field.show"></i>\
              {{field.label}}\
            </a>\
          </li>\
        </button-dropdown>\
        <button-dropdown :text="highlightMode.label" b-class="btn-xs" :dropup="true">\
          <li v-for="mode in highlightModes">\
            <a href="#" @click="highlightMode = mode">\
              <i class="fa fa-check text-success right"\
                 aria-hidden="true" v-show="highlightMode == mode"></i>\
              {{mode.label}}\
            </a>\
          </li>\
        </button-dropdown>\
        <button-state class="btn-xs right"\
                      v-model="autoRefresh"\
                      enabled-text="Auto refresh on"\
                      disabled-text="Auto refresh off">\
        </button-state>\
        <button-dropdown v-if="autoRefresh" :text="intervalText" class="right" b-class="btn-xs" :dropup="true">\
          <li><a href="#" @click="interval = 1000">1s</a></li>\
          <li><a href="#" @click="interval = 5000">5s</a></li>\
          <li><a href="#" @click="interval = 10000">10s</a></li>\
          <li><a href="#" @click="interval = 20000">20s</a></li>\
          <li><a href="#" @click="interval = 40000">40s</a></li>\
        </button-dropdown>\
        <button class="btn btn-default btn-xs right"\
                type="button"\
                @click="getFlows"\
                title="Refresh flows"\
                v-if="!autoRefresh">\
          <i class="fa fa-refresh" aria-hidden="true"></i>\
        </button>\
      </div>\
    </div>\
    <div v-else class="alert-danger">{{queryError}}</div>\
  ',

  components: {
    'table-header': TableHeader,
  },

  data: function() {
    return {
      queryResults: [],
      queryError: "",
      sortBy: null,
      sortOrder: -1,
      interval: 1000,
      intervalId: null,
      autoRefresh: false,
      showDetail: {},
      highlightModes: [
        {
          field: 'TrackingID',
          label: 'Follow L2',
        },
        {
          field: 'L3TrackingID',
          label: 'Follow L3',
        },
      ],
      highlightMode: null,
      fields: [
        {
          name: ['UUID'],
          label: 'UUID',
          show: false,
        },
        {
          name: ['LayersPath'],
          label: 'Layers',
          show: false,
        },
        {
          name: ['Application'],
          label: 'App.',
          show: true,
        },
        {
          name: ['Network.Protocol', 'Link.Protocol'],
          label: 'Proto.',
          show: false,
        },
        {
          name: ['Network.A', 'Link.A'],
          label: 'A',
          show: true,
        },
        {
          name: ['Network.B', 'Link.B'],
          label: 'B',
          show: true,
        },
        {
          name: ['Transport.Protocol'],
          label: 'L4 Proto.',
          show: false,
        },
        {
          name: ['Transport.A'],
          label: 'A port',
          show: false,
        },
        {
          name: ['Transport.B'],
          label: 'B port',
          show: false,
        },
        {
          name: ['Metric.ABPackets'],
          label: 'AB Pkts',
          show: true,
        },
        {
          name: ['Metric.BAPackets'],
          label: 'BA Pkts',
          show: true,
        },
        {
          name: ['Metric.ABBytes'], 
          label: 'AB Bytes',
          show: true,
        },
        {
          name: ['Metric.BABytes'],
          label: 'BA Bytes',
          show: true,
        },
        {
          name: ['TrackingID'],
          label: 'L2 Tracking ID',
          show: false,
        },
        {
          name: ['L3TrackingID'],
          label: 'L3 Tracking ID',
          show: false,
        },
        {
          name: ['NodeTID'],
          label: 'Interface',
          show: false,
        },
      ]
    };
  },

  created: function() {
    // sort by Application by default
    this.sortBy = this.fields[2].name;
    // use L2 mode by default
    this.highlightMode = this.highlightModes[0];
    this.getFlows();
  },

  beforeDestroy: function() {
    this.stopAutoRefresh();
  },

  watch: {

    autoRefresh: function(newVal) {
      if (newVal === true)
        this.startAutoRefresh();
      else
        this.stopAutoRefresh();
    },

    interval: function() {
      this.stopAutoRefresh();
      this.startAutoRefresh();
    },

    value: function() {
      this.getFlows();
    },

  },

  computed: {

    sortedResults: function() {
      return this.queryResults.sort(this.compareFlows);
    },

    // When Dedup() is used we show the detail of
    // the flow using TrackingID because the flow
    // returned has not always the same UUID
    showDetailField: function() {
      if (this.value.search('Dedup') !== -1) {
        return 'TrackingID';
      }
      return 'UUID';
    },

    intervalText: function() {
      return "Every " + this.interval / 1000 + "s";
    },

  },

  methods: {

    startAutoRefresh: function() {
      this.intervalId = setInterval(this.getFlows.bind(this), this.interval);
    },

    stopAutoRefresh: function() {
      if (this.intervalId !== null) {
        clearInterval(this.intervalId);
        this.intervalId = null;
      }
    },

    getFlows: function() {
      var self = this;
      TopologyAPI.query(this.value)
        .then(function(r) {
          self.queryResults = r; 
        })
        .fail(function(r) {
          self.queryError = r.responseText;
          self.stopAutoRefresh();
        });
    },

    // Keep track of which flow detail we should display
    showFlowDetail: function(obj) {
      if (this.showDetail[obj[this.showDetailField]]) { 
        Vue.delete(this.showDetail, obj[this.showDetailField]);
      } else {
        Vue.set(this.showDetail, obj[this.showDetailField], true);
      }
    },

    highlightNodes: function(obj, bool) {
      var query = "G.Flows().Has('" + this.highlightMode.field + "', '" + obj[this.highlightMode.field] + "').Hops()";
      TopologyAPI.query(query)
        .then(function(nodes) {
          nodes.forEach(function(n) {
            topologyLayout.SetNodeClass(n.ID, "highlighted", bool);
          });
        });
    },

    compareFlows: function(f1, f2) {
      if (!this.sortBy) {
        return 0;
      }
      var f1FieldValue = this.fieldValue(f1, this.sortBy),
          f2FieldValue = this.fieldValue(f2, this.sortBy);
      if (f1FieldValue < f2FieldValue)
        return -1 * this.sortOrder;
      if (f1FieldValue > f2FieldValue)
        return 1 * this.sortOrder;
      return 0;
    },

    fieldValue: function(object, paths) {
      for (var path of paths) {
        var value = object;
        for (var k of path.split(".")) {
          if (value[k] !== undefined) {
            value = value[k];
          } else {
            value = null;
            break;
          }
        }
        if (value !== null) {
          return value;
        }
      }
      return "";
    },

    sort: function(name) {
      if (this.sortBy == name) {
        this.sortOrder = this.sortOrder * -1;
      } else {
        this.sortOrder = -1;
      }
      this.sortBy = name;
    },

  },

});
