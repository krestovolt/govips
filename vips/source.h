#include <vips/vips.h>

extern gint64 goSourceRead(gpointer go_source, void *buffer, gint64 length);

extern gint64 goSourceSeek(gpointer go_source, gint64 offset, int whence);

VipsSourceCustom *create_go_custom_source();

gulong connect_go_signal_read(VipsSourceCustom *source_custom,
                              gpointer go_source);

gulong connect_go_signal_seek(VipsSourceCustom *source_custom,
                              gpointer go_source);

void free_go_custom_source(VipsSourceCustom *source_custom,
                           gulong rsig_handler_id, gulong ssig_handler_id);

static gint64 go_read(VipsSourceCustom *source_custom, void *buffer,
                      gint64 length, gpointer go_source);
static gint64 go_seek(VipsSourceCustom *source_custom, gint64 offset,
                      int whence, gpointer go_source);
