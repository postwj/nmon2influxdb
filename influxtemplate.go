// nmon2influx
// import nmon report in Influxdb
// author: adejoux@djouctech.net

package main

const influxtempl = `
{
  "id": null,
  "title": "{{.Hostname}} nmon report",
  "originalTitle": "{{.Hostname}} nmon report",
  "tags": [],
  "style": "dark",
  "timezone": "browser",
  "editable": true,
  "hideControls": false,
  "rows": [
      {
      "title": "INFORMATION",
      "height": "250px",
      "editable": true,
      "collapse": true,
      "panels": [
        {
          "error": false,
          "span": 12,
          "editable": true,
          "type": "text",
          "id": 10,
          "mode": "html",
          "content": "<table>{{.TextContent}}</table>",
          "style": {},
          "title": "INFORMATION"
        }
      ]
    },
    {
      "title": "CPU",
      "height": "250px",
      "editable": true,
      "collapse": false,
      "panels": [
        {
          "error": false,
          "span": 4,
          "editable": true,
          "type": "graph",
          "id": 1,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 1,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": true,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "individual",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "User%",
              "series": "{{$.Hostname}}_CPU_ALL",
              "query": "select mean(\"User%\") from \"{{$.Hostname}}_CPU_ALL\" where $timeFilter group by time($interval) order asc",
              "alias": "User%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "Sys%",
              "series": "{{$.Hostname}}_CPU_ALL",
              "query": "select mean(\"Sys%\") from \"{{$.Hostname}}_CPU_ALL\" where $timeFilter group by time($interval) order asc",
              "alias": "Sys%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "Wait%",
              "series": "{{$.Hostname}}_CPU_ALL",
              "query": "select mean(\"Wait%\") from \"{{$.Hostname}}_CPU_ALL\" where $timeFilter group by time($interval) order asc",
              "alias": "Wait%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "Idle%",
              "series": "{{$.Hostname}}_CPU_ALL",
              "query": "select mean(\"Idle%\") from \"{{$.Hostname}}_CPU_ALL\" where $timeFilter group by time($interval) order asc",
              "alias": "Idle%",
              "rawQuery": true,
              "hide": false
            }
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "CPU_ALL",
          "leftYAxisLabel": "%"
        },
        {
          "error": false,
          "span": 4,
          "editable": true,
          "type": "graph",
          "id": 2,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 1,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": true,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "EC_User%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"EC_User%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "EC_User%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "EC_Sys%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"EC_Sys%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "EC_Sys%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "EC_Wait%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"EC_Wait%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "EC_Wait%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "EC_Idle%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"EC_Idle%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "EC_Idle%",
              "rawQuery": true,
              "hide": false
            }
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "LPAR",
          "leftYAxisLabel": "%"
        },
        {
          "error": false,
          "span": 4,
          "editable": true,
          "type": "graph",
          "id": 2,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 1,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": true,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "VP_User%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"VP_User%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "VP_User%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "VP_Sys%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"VP_Sys%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "VP_Sys%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "VP_Wait%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"VP_Wait%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "VP_Wait%",
              "rawQuery": true,
              "hide": false
            },
            {
              "function": "mean",
              "column": "VP_Idle%",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(\"VP_Wait%\") from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "VP_Wait%",
              "rawQuery": true,
              "hide": false
            }
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "LPAR",
          "leftYAxisLabel": "%"
        }
      ]
    },
    {
      "title": "LPAR",
      "height": "250px",
      "editable": true,
      "collapse": false,
      "panels": [
        {
          "error": false,
          "span": 12,
          "editable": true,
          "type": "graph",
          "id": 6,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 0,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": false,
          "legend": {
            "show": true,
            "values": true,
            "min": true,
            "max": true,
            "current": false,
            "total": false,
            "avg": true
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "PhysicalCPU",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(PhysicalCPU) from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "PhysicalCPU"
            },
            {
              "function": "mean",
              "column": "entitled",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(entitled) from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "entitled"
            },
            {
              "function": "mean",
              "column": "Folded",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(Folded) from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "folded"
            },
            {
              "function": "mean",
              "column": "virtualCPUs",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(virtualCPUs) from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "virtualCPUs"
            },
            {
              "function": "mean",
              "column": "logicalCPUs",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(logicalCPUs) from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "logicalCPUs"
            },
            {
              "function": "mean",
              "column": "PoolIdle",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(PoolIdle) from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "PoolIdle"
            },
            {
              "function": "mean",
              "column": "poolCPUs",
              "series": "{{$.Hostname}}_LPAR",
              "query": "select mean(poolCPUs) from \"{{$.Hostname}}_LPAR\" where $timeFilter group by time($interval) order asc",
              "alias": "poolCPUs"
            }
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "LPAR",
          "leftYAxisLabel": "cpu"
        }
      ]
    },
    {
      "title": "MEM",
      "height": "250px",
      "editable": true,
      "collapse": true,
      "panels": [
        {
          "error": false,
          "span": 4,
          "editable": true,
          "type": "graph",
          "id": 4,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 0,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": false,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "%minperm",
              "series": "{{$.Hostname}}_MEMUSE",
              "query": "select mean(\"%minperm\") from \"{{$.Hostname}}_MEMUSE\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "%minperm"
            },
            {
              "function": "mean",
              "column": "%minperm",
              "series": "{{$.Hostname}}_MEMUSE",
              "query": "select mean(\"%maxperm\") from \"{{$.Hostname}}_MEMUSE\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "%maxperm"
            },
            {
              "function": "mean",
              "column": "%numperm",
              "series": "{{$.Hostname}}_MEMUSE",
              "query": "select mean(\"%numperm\") from \"{{$.Hostname}}_MEMUSE\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "%numperm"
            },
            {
              "function": "mean",
              "column": "%numperm",
              "series": "{{$.Hostname}}_MEMUSE",
              "query": "select mean(\"%numclient\") from \"{{$.Hostname}}_MEMUSE\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "%numclient"
            }
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "MEMUSE",
          "leftYAxisLabel": "%"
        },
        {
          "error": false,
          "span": 4,
          "editable": true,
          "type": "graph",
          "id": 5,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 1,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": false,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "Real free(MB)",
              "series": "{{$.Hostname}}_MEM",
              "query": "select mean(\"Real free(MB)\") from \"{{$.Hostname}}_MEM\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "Real free(MB)"
            },
            {
              "function": "mean",
              "column": "Virtual free(MB)",
              "series": "{{$.Hostname}}_MEM",
              "query": "select mean(\"Virtual free(MB)\") from \"{{$.Hostname}}_MEM\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "Virtual free(MB)"
            },
            {
              "function": "mean",
              "column": "Real total(MB)",
              "series": "{{$.Hostname}}_MEM",
              "query": "select mean(\"Real total(MB)\") from \"{{$.Hostname}}_MEM\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "Real total(MB)"
            },
            {
              "function": "mean",
              "column": "Virtual total(MB)",
              "series": "{{$.Hostname}}_MEM",
              "query": "select mean(\"Virtual total(MB)\") from \"{{$.Hostname}}_MEM\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "Virtual total(MB)"
            }
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "MEM",
          "leftYAxisLabel": "MB"
        },
        {
          "error": false,
          "span": 4,
          "editable": true,
          "type": "graph",
          "id": 5,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 1,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": true,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "Process%",
              "series": "{{$.Hostname}}_MEMNEW",
              "query": "select mean(\"Process%\") from \"{{$.Hostname}}_MEMNEW\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "Process%"
            },
            {
              "function": "mean",
              "column": "FScache%",
              "series": "{{$.Hostname}}_MEMNEW",
              "query": "select mean(\"FScache%\") from \"{{$.Hostname}}_MEMNEW\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "FScache%"
            },
            {
              "function": "mean",
              "column": "System%",
              "series": "{{$.Hostname}}_MEMNEW",
              "query": "select mean(\"System%\") from \"{{$.Hostname}}_MEMNEW\" where $timeFilter group by time($interval) order asc",
              "rawQuery": true,
              "alias": "System%"
            }

          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "MEMNEW",
          "leftYAxisLabel": "MB"
        }
      ]
    },
    {
      "title": "IOADAPT",
      "height": "250px",
      "editable": true,
      "collapse": true,
      "panels": [
        {
          "error": false,
          "span": 12,
          "editable": true,
          "type": "graph",
          "id": 3,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": false,
          "fill": 0,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": true,
          "stack": true,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "individual",
            "query_as_alias": true
          },
          "targets": [
            {{ range $index, $adapter := .GetColumns "IOADAPT" }}{{ if $index}},{{end}}
                {
                  "function": "mean",
                  "column": "{{$adapter}}",
                  "series": "{{$.Hostname}}_IOADAPT",
                  "query": "select mean(\"{{$adapter}}\") from \"{{$.Hostname}}_IOADAPT\" where $timeFilter group by time($interval) order asc",
                  "rawQuery": true,
                  "alias": "{{$adapter}}"
                }{{end}}
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "IOADAPT",
          "leftYAxisLabel": "KB/s"
        }
      ]
    },
    {{ if .GetColumns "TEST" }}
        titi
    {{end}}
    {
      "title": "NET",
      "height": "250px",
      "editable": true,
      "collapse": true,
      "panels": [
        {
          "error": false,
          "span": 12,
          "editable": true,
          "type": "graph",
          "id": 7,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 1,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": true,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {{ range $index, $adapter := .GetColumns "NET" }}{{ if $index}},{{end}}
                {
                    "function": "mean",
                    "column": "{{$adapter}}",
                    "series": "{{$.Hostname}}_NET",
                    "query": "select mean(\"{{$adapter}}\") from \"{{$.Hostname}}_NET\" where $timeFilter group by time($interval) order asc",
                    "rawQuery": true,
                    "alias": "{{$adapter}}"
                }{{end}}
            ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "NET"
        }
      ]
    },
    {
      "title": "PAGE",
      "height": "250px",
      "editable": true,
      "collapse": true,
      "panels": [
        {
          "error": false,
          "span": 12,
          "editable": true,
          "type": "graph",
          "id": 8,
          "datasource": null,
          "renderer": "flot",
          "x-axis": true,
          "y-axis": true,
          "scale": 1,
          "y_formats": [
            "short",
            "short"
          ],
          "grid": {
            "leftMax": null,
            "rightMax": null,
            "leftMin": null,
            "rightMin": null,
            "threshold1": null,
            "threshold2": null,
            "threshold1Color": "rgba(216, 200, 27, 0.27)",
            "threshold2Color": "rgba(234, 112, 112, 0.22)"
          },
          "annotate": {
            "enable": false
          },
          "resolution": 100,
          "lines": true,
          "fill": 0,
          "linewidth": 1,
          "points": false,
          "pointradius": 5,
          "bars": false,
          "stack": false,
          "legend": {
            "show": true,
            "values": false,
            "min": false,
            "max": false,
            "current": false,
            "total": false,
            "avg": false
          },
          "percentage": false,
          "zerofill": true,
          "nullPointMode": "connected",
          "steppedLine": false,
          "tooltip": {
            "value_type": "cumulative",
            "query_as_alias": true
          },
          "targets": [
            {
              "function": "mean",
              "column": "pgin",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(pgin) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "pgin"
            },
            {
              "function": "mean",
              "column": "pgout",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(pgout) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "pgout"
            },
            {
              "function": "mean",
              "column": "scans",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(scans) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "scans"
            },
            {
              "function": "mean",
              "column": "cycles",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(cycles) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "cycles"
            },
            {
              "function": "mean",
              "column": "faults",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(faults) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "faults"
            },
            {
              "function": "mean",
              "column": "pgsin",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(pgsin) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "pgsin"
            },
            {
              "function": "mean",
              "column": "pgsout",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(pgsout) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "pgsout"
            },
            {
              "function": "mean",
              "column": "reclaims",
              "series": "{{$.Hostname}}_PAGE",
              "query": "select mean(reclaims) from \"{{$.Hostname}}_PAGE\" where $timeFilter group by time($interval) order asc",
              "alias": "reclaims"
            }
          ],
          "aliasColors": {},
          "seriesOverrides": [],
          "title": "PAGE"
        }
      ]
    }
  ],
  "nav": [
    {
      "type": "timepicker",
      "enable": true,
      "status": "Stable",
      "time_options": [
        "5m",
        "15m",
        "1h",
        "6h",
        "12h",
        "24h",
        "2d",
        "7d",
        "30d"
      ],
      "refresh_intervals": [
        "5s",
        "10s",
        "30s",
        "1m",
        "5m",
        "15m",
        "30m",
        "1h",
        "2h",
        "1d"
      ],
      "now": false,
      "collapse": false,
      "notice": false
    }
  ],
  "time": {
    "from": "2014-11-10T00:44:32.089Z",
    "to": "2014-11-13T16:37:30.528Z",
    "now": false
  },
  "templating": {
    "list": []
  },
  "annotations": {
    "list": [],
    "enable": false
  },
  "refresh": false,
  "version": 6
}
`






