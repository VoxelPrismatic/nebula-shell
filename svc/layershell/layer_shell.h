#ifndef LAYER_SHELL_H
#define LAYER_SHELL_H
#include <stdbool.h>
#include <stdint.h>

void UseLayerShell();
void WinSetAnchors(void *window, int anchors);
int WinGetAnchors(void *window);
void WinSetExclusionZone(void *window, int32_t zone);
int32_t WinGetExclusionZone(void *window);
void WinSetExclusiveEdge(void *window, int edge);
int WinGetExclusiveEdge(void *window);
void WinSetMargins(void *window, void *margins);
void *WinGetMargins(void *window);
void WinSetDesiredSize(void *window, void *size);
void *WinGetDesiredSize(void *window);
void WinSetKeyboardInteractivity(void *window, int kbd);
int WinGetKeyboardInteractivity(void *window);
void WinSetLayer(void *window, int layer);
int WinGetLayer(void *window);
void WinSetScreenConfiguration(void *window, int config);
int WinGetScreenConfiguration(void *window);
void WinSetScope(void *window, char *str);
char *WinGetScope(void *window);
void FreeBuf(void *buf);
void WinSetCloseOnDismissed(void *window, bool close);
bool WinGetCloseOnDismissed(void *window);

#endif // LAYER_SHELL_H
